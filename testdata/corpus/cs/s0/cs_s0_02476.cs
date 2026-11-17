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

    public struct Point23120
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid23120
    {
        public static Point23120 Make(int left, int right)
        {
            return new Point23120 { X = left, Y = right };
        }
    }
    public sealed class SlotMap23121
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
    public static int Scale23122<T>(T value) where T : struct
    {
        int number = 5;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
}
