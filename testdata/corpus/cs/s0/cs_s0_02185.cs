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

    public delegate void Notifier20400(int value);

    public sealed class Source20400
    {
        public event Notifier20400 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap20401
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
    public sealed class SlotMap20402
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
    public interface IHandler20403
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20403(string Label, int Payload);

    public sealed class Handler20403 : IHandler20403
    {
        public void Handle(int value) { }
    }
    public static int Scale20404<T>(T value) where T : struct
    {
        int number = 72;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
}
