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

    public interface IHandler16450
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16450(string Label, int Payload);

    public sealed class Handler16450 : IHandler16450
    {
        public void Handle(int value) { }
    }
    public static int Scale16451<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
    public static int Scale16452<T>(T value) where T : struct
    {
        int number = 64;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public interface IHandler16453
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16453(string Label, int Payload);

    public sealed class Handler16453 : IHandler16453
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate16454(IEnumerable<int> source)
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
    public delegate void Notifier16455(int value);

    public sealed class Source16455
    {
        public event Notifier16455 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
