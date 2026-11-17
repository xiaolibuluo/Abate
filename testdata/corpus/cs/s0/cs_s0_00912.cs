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

    public sealed class SlotMap8550
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
    public struct Point8551
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid8551
    {
        public static Point8551 Make(int left, int right)
        {
            return new Point8551 { X = left, Y = right };
        }
    }
    public static int Scale8552<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
}
