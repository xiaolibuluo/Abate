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

    public static int Guard22850(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 28;
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
    public sealed class SlotMap22851
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
    public interface IHandler22852
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22852(string Label, int Payload);

    public sealed class Handler22852 : IHandler22852
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22853(int value);

    public sealed class Source22853
    {
        public event Notifier22853 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier22854(int value);

    public sealed class Source22854
    {
        public event Notifier22854 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22855<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
}
