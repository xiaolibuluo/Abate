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

    public struct Point12390
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid12390
    {
        public static Point12390 Make(int left, int right)
        {
            return new Point12390 { X = left, Y = right };
        }
    }
    public struct Point12391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 17;
    }

    public static class Grid12391
    {
        public static Point12391 Make(int left, int right)
        {
            return new Point12391 { X = left, Y = right };
        }
    }
    public sealed class SlotMap12392
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
