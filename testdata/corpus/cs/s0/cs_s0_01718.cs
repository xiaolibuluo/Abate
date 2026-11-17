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

    public static int Scale16040<T>(T value) where T : struct
    {
        int number = 99;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public static int Scale16041<T>(T value) where T : struct
    {
        int number = 9;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public delegate void Notifier16042(int value);

    public sealed class Source16042
    {
        public event Notifier16042 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale16043<T>(T value) where T : struct
    {
        int number = 23;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public delegate void Notifier16044(int value);

    public sealed class Source16044
    {
        public event Notifier16044 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler16045
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16045(string Label, int Payload);

    public sealed class Handler16045 : IHandler16045
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap16046
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
