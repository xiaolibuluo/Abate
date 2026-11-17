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

    public struct Point8000
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid8000
    {
        public static Point8000 Make(int left, int right)
        {
            return new Point8000 { X = left, Y = right };
        }
    }
    public struct Point8001
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid8001
    {
        public static Point8001 Make(int left, int right)
        {
            return new Point8001 { X = left, Y = right };
        }
    }
    public sealed class SlotMap8002
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
