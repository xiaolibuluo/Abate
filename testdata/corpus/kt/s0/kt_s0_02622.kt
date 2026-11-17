package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21640<A, B>(val first: A, val second: B)

fun combine21640(left: Int, right: Int): Pair21640<Int, Int> {
    return Pair21640(left + 7, right + 29)
}
fun <T> List<T>.secondOrNull21641(): T? = if (size > 1) this[1] else null

fun render21641(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21642<A, B>(val first: A, val second: B)

fun combine21642(left: Int, right: Int): Pair21642<Int, Int> {
    return Pair21642(left + 21, right + 51)
}
class Counter21643 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome21644 {
    data class Success(val payload: String) : Outcome21644()
    data object Failure : Outcome21644()
}

fun describe21644(outcome: Outcome21644): String = when (outcome) {
    is Outcome21644.Success -> outcome.payload
    Outcome21644.Failure -> "failed"
}
