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

    public delegate void Notifier14370(int value);

    public sealed class Source14370
    {
        public event Notifier14370 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier14371(int value);

    public sealed class Source14371
    {
        public event Notifier14371 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14372
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14372(string Label, int Payload);

    public sealed class Handler14372 : IHandler14372
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate14373(IEnumerable<int> source)
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
