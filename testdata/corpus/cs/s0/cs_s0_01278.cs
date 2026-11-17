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

    public delegate void Notifier11980(int value);

    public sealed class Source11980
    {
        public event Notifier11980 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler11981
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11981(string Label, int Payload);

    public sealed class Handler11981 : IHandler11981
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11982(int value);

    public sealed class Source11982
    {
        public event Notifier11982 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate11983(IEnumerable<int> source)
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
}
