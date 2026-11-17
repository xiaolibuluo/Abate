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

    public struct Point21270
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid21270
    {
        public static Point21270 Make(int left, int right)
        {
            return new Point21270 { X = left, Y = right };
        }
    }
    public interface IHandler21271
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21271(string Label, int Payload);

    public sealed class Handler21271 : IHandler21271
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier21272(int value);

    public sealed class Source21272
    {
        public event Notifier21272 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap21273
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
