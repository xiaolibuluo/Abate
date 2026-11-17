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

    public static int Scale19320<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public struct Point19321
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid19321
    {
        public static Point19321 Make(int left, int right)
        {
            return new Point19321 { X = left, Y = right };
        }
    }
    public sealed class SlotMap19322
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
    public struct Point19323
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 19;
    }

    public static class Grid19323
    {
        public static Point19323 Make(int left, int right)
        {
            return new Point19323 { X = left, Y = right };
        }
    }
}
