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

    public delegate void Notifier11050(int value);

    public sealed class Source11050
    {
        public event Notifier11050 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler11051
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11051(string Label, int Payload);

    public sealed class Handler11051 : IHandler11051
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap11052
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
    public delegate void Notifier11053(int value);

    public sealed class Source11053
    {
        public event Notifier11053 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale11054<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
}
