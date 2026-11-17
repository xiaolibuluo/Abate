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

    public delegate void Notifier7130(int value);

    public sealed class Source7130
    {
        public event Notifier7130 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler7131
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7131(string Label, int Payload);

    public sealed class Handler7131 : IHandler7131
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap7132
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
    public static int Scale7133<T>(T value) where T : struct
    {
        int number = 80;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
}
