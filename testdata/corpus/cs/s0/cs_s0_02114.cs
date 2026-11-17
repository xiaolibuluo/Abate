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

    public sealed class SlotMap19730
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
    public interface IHandler19731
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19731(string Label, int Payload);

    public sealed class Handler19731 : IHandler19731
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier19732(int value);

    public sealed class Source19732
    {
        public event Notifier19732 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
