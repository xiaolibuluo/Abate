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

    public delegate void Notifier23250(int value);

    public sealed class Source23250
    {
        public event Notifier23250 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier23251(int value);

    public sealed class Source23251
    {
        public event Notifier23251 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23252
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23252(string Label, int Payload);

    public sealed class Handler23252 : IHandler23252
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate23253(IEnumerable<int> source)
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
    public delegate void Notifier23254(int value);

    public sealed class Source23254
    {
        public event Notifier23254 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap23255
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
