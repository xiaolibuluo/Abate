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

    public delegate void Notifier23270(int value);

    public sealed class Source23270
    {
        public event Notifier23270 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap23271
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
    public static int Scale23272<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 33; }
        return number;
    }
    public delegate void Notifier23273(int value);

    public sealed class Source23273
    {
        public event Notifier23273 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23274
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23274(string Label, int Payload);

    public sealed class Handler23274 : IHandler23274
    {
        public void Handle(int value) { }
    }
    public static int Scale23275<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
}
