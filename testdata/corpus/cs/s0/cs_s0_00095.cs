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

    public static int Scale920<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public interface IHandler921
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope921(string Label, int Payload);

    public sealed class Handler921 : IHandler921
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap922
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
    public delegate void Notifier923(int value);

    public sealed class Source923
    {
        public event Notifier923 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler924
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope924(string Label, int Payload);

    public sealed class Handler924 : IHandler924
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier925(int value);

    public sealed class Source925
    {
        public event Notifier925 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
