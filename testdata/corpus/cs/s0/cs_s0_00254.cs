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

    public struct Point2410
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid2410
    {
        public static Point2410 Make(int left, int right)
        {
            return new Point2410 { X = left, Y = right };
        }
    }
    public sealed class SlotMap2411
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
    public sealed class SlotMap2412
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
    public interface IHandler2413
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2413(string Label, int Payload);

    public sealed class Handler2413 : IHandler2413
    {
        public void Handle(int value) { }
    }
    public interface IHandler2414
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2414(string Label, int Payload);

    public sealed class Handler2414 : IHandler2414
    {
        public void Handle(int value) { }
    }
}
