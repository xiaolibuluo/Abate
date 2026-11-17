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

    public delegate void Notifier5230(int value);

    public sealed class Source5230
    {
        public event Notifier5230 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5231
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5231(string Label, int Payload);

    public sealed class Handler5231 : IHandler5231
    {
        public void Handle(int value) { }
    }
    public interface IHandler5232
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5232(string Label, int Payload);

    public sealed class Handler5232 : IHandler5232
    {
        public void Handle(int value) { }
    }
    public struct Point5233
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid5233
    {
        public static Point5233 Make(int left, int right)
        {
            return new Point5233 { X = left, Y = right };
        }
    }
    public sealed class SlotMap5234
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
