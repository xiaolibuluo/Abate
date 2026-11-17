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

    public async Task<int> Aggregate7240(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public interface IHandler7241
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7241(string Label, int Payload);

    public sealed class Handler7241 : IHandler7241
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier7242(int value);

    public sealed class Source7242
    {
        public event Notifier7242 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier7243(int value);

    public sealed class Source7243
    {
        public event Notifier7243 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard7244(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 36;
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
    public sealed class SlotMap7245
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
