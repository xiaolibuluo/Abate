/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// csRecipe covers namespaces, records, generics with constraints, properties,
// indexers, events, LINQ, async/await, attributes and checked arithmetic.
func csRecipe() Recipe {
	return Recipe{
		Language: "cs",
		Ext:      "cs",
		Tags:     []string{"class", "generic", "linq", "async", "exception", "attribute", "property", "event"},
		Header: `using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
`,
		Footer: "}\n",
		AnchorClean: `    public static bool AnchorProbe()
    {
        return 1 + 2 == 3;
    }

`,
		AnchorFull: `    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

`,
		CleanBlock: "    public static int Step@@I@@x@@S@@() { return @@N@@ + @@S@@; }\n",
		Blocks: []BlockTemplate{
			{Text: `    public interface IHandler@@I@@@@S@@
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope@@I@@@@S@@(string Label, int Payload);

    public sealed class Handler@@I@@@@S@@ : IHandler@@I@@@@S@@
    {
        public void Handle(int value) { }
    }
`},
			{Text: `    public async Task<int> Aggregate@@I@@@@S@@(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
`},
			{Text: `    public sealed class SlotMap@@I@@@@S@@
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
`},
			{Text: `    public delegate void Notifier@@I@@@@S@@(int value);

    public sealed class Source@@I@@@@S@@
    {
        public event Notifier@@I@@@@S@@ Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
`},
			{Text: `    public static int Scale@@I@@@@S@@<T>(T value) where T : struct
    {
        int number = @@N@@;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + @@M@@; }
        return number;
    }
`},
			{Text: `    public static int Guard@@I@@@@S@@(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + @@N@@;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
`},
			{Text: `    public struct Point@@I@@@@S@@
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + @@N@@;
    }

    public static class Grid@@I@@@@S@@
    {
        public static Point@@I@@@@S@@ Make(int left, int right)
        {
            return new Point@@I@@@@S@@ { X = left, Y = right };
        }
    }
`},
		},
	}
}
