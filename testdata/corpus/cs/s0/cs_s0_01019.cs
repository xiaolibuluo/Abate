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

    public async Task<int> Aggregate9550(IEnumerable<int> source)
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
    public interface IHandler9551
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9551(string Label, int Payload);

    public sealed class Handler9551 : IHandler9551
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier9552(int value);

    public sealed class Source9552
    {
        public event Notifier9552 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard9553(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
