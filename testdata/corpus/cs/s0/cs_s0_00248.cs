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

    public delegate void Notifier2350(int value);

    public sealed class Source2350
    {
        public event Notifier2350 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler2351
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2351(string Label, int Payload);

    public sealed class Handler2351 : IHandler2351
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap2352
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
    public struct Point2353
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid2353
    {
        public static Point2353 Make(int left, int right)
        {
            return new Point2353 { X = left, Y = right };
        }
    }
    public struct Point2354
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid2354
    {
        public static Point2354 Make(int left, int right)
        {
            return new Point2354 { X = left, Y = right };
        }
    }
}
