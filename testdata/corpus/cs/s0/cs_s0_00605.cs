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

    public delegate void Notifier5660(int value);

    public sealed class Source5660
    {
        public event Notifier5660 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5661
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5661(string Label, int Payload);

    public sealed class Handler5661 : IHandler5661
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap5662
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
    public static int Guard5663(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 12;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
