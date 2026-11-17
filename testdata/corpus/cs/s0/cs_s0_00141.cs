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

    public static int Scale1350<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public delegate void Notifier1351(int value);

    public sealed class Source1351
    {
        public event Notifier1351 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale1352<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public interface IHandler1353
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1353(string Label, int Payload);

    public sealed class Handler1353 : IHandler1353
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap1354
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
