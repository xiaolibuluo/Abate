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

    public interface IHandler24570
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24570(string Label, int Payload);

    public sealed class Handler24570 : IHandler24570
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier24571(int value);

    public sealed class Source24571
    {
        public event Notifier24571 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate24572(IEnumerable<int> source)
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
