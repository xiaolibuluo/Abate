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

    public struct Point23170
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid23170
    {
        public static Point23170 Make(int left, int right)
        {
            return new Point23170 { X = left, Y = right };
        }
    }
    public struct Point23171
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid23171
    {
        public static Point23171 Make(int left, int right)
        {
            return new Point23171 { X = left, Y = right };
        }
    }
    public sealed class SlotMap23172
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
