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

    public struct Point2710
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid2710
    {
        public static Point2710 Make(int left, int right)
        {
            return new Point2710 { X = left, Y = right };
        }
    }
    public interface IHandler2711
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2711(string Label, int Payload);

    public sealed class Handler2711 : IHandler2711
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap2712
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
