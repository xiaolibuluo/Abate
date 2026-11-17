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

    public sealed class SlotMap10090
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
    public struct Point10091
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid10091
    {
        public static Point10091 Make(int left, int right)
        {
            return new Point10091 { X = left, Y = right };
        }
    }
    public interface IHandler10092
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10092(string Label, int Payload);

    public sealed class Handler10092 : IHandler10092
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier10093(int value);

    public sealed class Source10093
    {
        public event Notifier10093 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
