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

    public static int Scale14840<T>(T value) where T : struct
    {
        int number = 91;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public struct Point14841
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid14841
    {
        public static Point14841 Make(int left, int right)
        {
            return new Point14841 { X = left, Y = right };
        }
    }
    public sealed class SlotMap14842
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
}
