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

    public sealed class SlotMap21720
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
    public interface IHandler21721
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21721(string Label, int Payload);

    public sealed class Handler21721 : IHandler21721
    {
        public void Handle(int value) { }
    }
    public static int Scale21722<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public delegate void Notifier21723(int value);

    public sealed class Source21723
    {
        public event Notifier21723 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
