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

    public interface IHandler23260
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23260(string Label, int Payload);

    public sealed class Handler23260 : IHandler23260
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier23261(int value);

    public sealed class Source23261
    {
        public event Notifier23261 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23262
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23262(string Label, int Payload);

    public sealed class Handler23262 : IHandler23262
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap23263
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
    public static int Guard23264(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 7;
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
