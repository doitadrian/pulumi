// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll
package dotnet

import (
	"strings"
	"text/template"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// nolint:lll
const csharpUtilitiesTemplateText = `// *** WARNING: this file was generated by {{.Tool}}. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Diagnostics.CodeAnalysis;
using System.IO;
using System.Reflection;
using Pulumi;

namespace {{.Namespace}}
{
    static class {{.ClassName}}
    {
        public static string? GetEnv(params string[] names)
        {
            foreach (var n in names)
            {
                var value = Environment.GetEnvironmentVariable(n);
                if (value != null)
                {
                    return value;
                }
            }
            return null;
        }

        static string[] trueValues = { "1", "t", "T", "true", "TRUE", "True" };
        static string[] falseValues = { "0", "f", "F", "false", "FALSE", "False" };
        public static bool? GetEnvBoolean(params string[] names)
        {
            var s = GetEnv(names);
            if (s != null)
            {
                if (Array.IndexOf(trueValues, s) != -1)
                {
                    return true;
                }
                if (Array.IndexOf(falseValues, s) != -1)
                {
                    return false;
                }
            }
            return null;
        }

        public static int? GetEnvInt32(params string[] names) => int.TryParse(GetEnv(names), out int v) ? (int?)v : null;

        public static double? GetEnvDouble(params string[] names) => double.TryParse(GetEnv(names), out double v) ? (double?)v : null;

        public static InvokeOptions WithVersion(this InvokeOptions? options)
        {
            if (options?.Version != null)
            {
                return options;
            }
            return new InvokeOptions
            {
                Parent = options?.Parent,
                Provider = options?.Provider,
                Version = Version,
            };
        }

        public static List<T> ToList<T>(this IEnumerable<T> elements)
        {
            return new List<T>(elements);
        }

        public static Input<List<T>> ToList<T>(this InputList<T> inputList)
        {
            return inputList.Apply(v => v.ToList());
        }

        public class Boxed
        {
            [AllowNull]
            public Object Value { get; }

            public Boxed([AllowNull] Object value)
            {
                Value = value;
            }

            public void Set(Object target, string propertyName)
            {
                var v = this.Value;
                if (v != null)
                {
                    var p = target.GetType().GetProperty(propertyName);
                    if (p != null)
                    {
                        p.SetValue(target, v);
                    }
                }
            }
        }

        public static Output<Boxed> Box<T>([AllowNull] this Input<T> input)
        {
            if (input == null)
            {
                return Output.Create(new Boxed(null));
            }
            else
            {
                return input.Apply(v => new Boxed(v));
            }
        }

        private readonly static string version;
        public static string Version => version;

        static Utilities()
        {
            var assembly = typeof(Utilities).GetTypeInfo().Assembly;
            using var stream = assembly.GetManifestResourceStream("{{.Namespace}}.version.txt");
            using var reader = new StreamReader(stream ?? throw new NotSupportedException("Missing embedded version.txt file"));
            version = reader.ReadToEnd().Trim();
            var parts = version.Split("\n");
            if (parts.Length == 2)
            {
                // The first part is the provider name.
                version = parts[1].Trim();
            }
        }
    }

    internal sealed class {{.Name}}ResourceTypeAttribute : Pulumi.ResourceTypeAttribute
    {
        public {{.Name}}ResourceTypeAttribute(string type) : base(type, Utilities.Version)
        {
        }
    }
}
`

var csharpUtilitiesTemplate = template.Must(template.New("CSharpUtilities").Parse(csharpUtilitiesTemplateText))

type csharpUtilitiesTemplateContext struct {
	Name      string
	Namespace string
	ClassName string
	Tool      string
}

// TODO(pdg): parameterize package name
const csharpProjectFileTemplateText = `<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <GeneratePackageOnBuild>true</GeneratePackageOnBuild>
    <Authors>Pulumi Corp.</Authors>
    <Company>Pulumi Corp.</Company>
    <Description>{{.Package.Description}}</Description>
    <PackageLicenseExpression>{{.Package.License}}</PackageLicenseExpression>
    <PackageProjectUrl>{{.Package.Homepage}}</PackageProjectUrl>
    <RepositoryUrl>{{.Package.Repository}}</RepositoryUrl>
    <PackageIcon>logo.png</PackageIcon>

    <TargetFramework>netcoreapp3.1</TargetFramework>
    <Nullable>enable</Nullable>
  </PropertyGroup>

  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|AnyCPU'">
    <GenerateDocumentationFile>true</GenerateDocumentationFile>
    <NoWarn>1701;1702;1591</NoWarn>
  </PropertyGroup>

  <PropertyGroup>
    <AllowedOutputExtensionsInPackageBuildOutputFolder>$(AllowedOutputExtensionsInPackageBuildOutputFolder);.pdb</AllowedOutputExtensionsInPackageBuildOutputFolder>
    <EmbedUntrackedSources>true</EmbedUntrackedSources>
    <PublishRepositoryUrl>true</PublishRepositoryUrl>
  </PropertyGroup>

  <PropertyGroup Condition="'$(GITHUB_ACTIONS)' == 'true'">
    <ContinuousIntegrationBuild>true</ContinuousIntegrationBuild>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="Microsoft.SourceLink.GitHub" Version="1.0.0" PrivateAssets="All" />
  </ItemGroup>

  <ItemGroup>
    <EmbeddedResource Include="version.txt" />
    <None Include="version.txt" Pack="True" PackagePath="content" />
  </ItemGroup>

  <ItemGroup>
    {{- range $package, $version := .PackageReferences}}
    <PackageReference Include="{{$package}}" Version="{{$version}}"{{if ispulumipkg $package}} ExcludeAssets="contentFiles"{{end}} />
    {{- end}}
  </ItemGroup>

  <ItemGroup>
    <None Include="logo.png">
      <Pack>True</Pack>
      <PackagePath></PackagePath>
    </None>
  </ItemGroup>

</Project>
`

var csharpProjectFileTemplate = template.Must(template.New("CSharpProject").Funcs(template.FuncMap{
	// ispulumipkg is used in the template to conditionally emit `ExcludeAssets="contentFiles"`
	// for `<PackageReference>`s that start with "Pulumi.", to prevent the references's contentFiles
	// from being included in this project's package. Otherwise, if a reference has version.txt
	// in its contentFiles, and we don't exclude contentFiles for the reference, the reference's
	// version.txt will be used over this project's version.txt.
	"ispulumipkg": func(s string) bool {
		return strings.HasPrefix(s, "Pulumi.")
	},
}).Parse(csharpProjectFileTemplateText))

type csharpProjectFileTemplateContext struct {
	XMLDoc            string
	Package           *schema.Package
	PackageReferences map[string]string
	Version           string
}
