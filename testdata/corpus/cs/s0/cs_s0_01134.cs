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

    public sealed class SlotMap10630
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
    public static int Scale10631<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public struct Point10632
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid10632
    {
        public static Point10632 Make(int left, int right)
        {
            return new Point10632 { X = left, Y = right };
        }
    }
    public struct Point10633
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid10633
    {
        public static Point10633 Make(int left, int right)
        {
            return new Point10633 { X = left, Y = right };
        }
    }
    public struct Point10634
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid10634
    {
        public static Point10634 Make(int left, int right)
        {
            return new Point10634 { X = left, Y = right };
        }
    }
}
