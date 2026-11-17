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

    public sealed class SlotMap2870
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
    public interface IHandler2871
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2871(string Label, int Payload);

    public sealed class Handler2871 : IHandler2871
    {
        public void Handle(int value) { }
    }
    public struct Point2872
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid2872
    {
        public static Point2872 Make(int left, int right)
        {
            return new Point2872 { X = left, Y = right };
        }
    }
    public sealed class SlotMap2873
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
