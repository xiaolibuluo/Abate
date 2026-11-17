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

    public struct Point13080
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid13080
    {
        public static Point13080 Make(int left, int right)
        {
            return new Point13080 { X = left, Y = right };
        }
    }
    public sealed class SlotMap13081
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
    public interface IHandler13082
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13082(string Label, int Payload);

    public sealed class Handler13082 : IHandler13082
    {
        public void Handle(int value) { }
    }
    public interface IHandler13083
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13083(string Label, int Payload);

    public sealed class Handler13083 : IHandler13083
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13084(int value);

    public sealed class Source13084
    {
        public event Notifier13084 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
