using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public struct Point4360
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 47;
    }

    public static class Grid4360
    {
        public static Point4360 Make(int left, int right)
        {
            return new Point4360 { X = left, Y = right };
        }
    }
    public sealed class SlotMap4361
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
    public struct Point4362
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid4362
    {
        public static Point4362 Make(int left, int right)
        {
            return new Point4362 { X = left, Y = right };
        }
    }
    public static int Guard4363(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
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
    public async Task<int> Aggregate4364(IEnumerable<int> source)
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
}
