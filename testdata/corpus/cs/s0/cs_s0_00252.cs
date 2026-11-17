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

    public static int Scale2390<T>(T value) where T : struct
    {
        int number = 8;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
    public struct Point2391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid2391
    {
        public static Point2391 Make(int left, int right)
        {
            return new Point2391 { X = left, Y = right };
        }
    }
    public static int Scale2392<T>(T value) where T : struct
    {
        int number = 22;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public sealed class SlotMap2393
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
    public async Task<int> Aggregate2394(IEnumerable<int> source)
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
