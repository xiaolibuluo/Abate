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

    public static int Scale20260<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public struct Point20261
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid20261
    {
        public static Point20261 Make(int left, int right)
        {
            return new Point20261 { X = left, Y = right };
        }
    }
    public sealed class SlotMap20262
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
    public sealed class SlotMap20263
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
