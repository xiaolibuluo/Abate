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

    public async Task<int> Aggregate10160(IEnumerable<int> source)
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
    public interface IHandler10161
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10161(string Label, int Payload);

    public sealed class Handler10161 : IHandler10161
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap10162
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
    public delegate void Notifier10163(int value);

    public sealed class Source10163
    {
        public event Notifier10163 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10164
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10164(string Label, int Payload);

    public sealed class Handler10164 : IHandler10164
    {
        public void Handle(int value) { }
    }
}
