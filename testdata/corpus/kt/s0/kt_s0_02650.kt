package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21890(): T? = if (size > 1) this[1] else null

fun render21890(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter21891 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
@JvmName("step21892")
fun step21892(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
data class Pair21893<A, B>(val first: A, val second: B)

fun combine21893(left: Int, right: Int): Pair21893<Int, Int> {
    return Pair21893(left + 62, right + 54)
}
data class Pair21894<A, B>(val first: A, val second: B)

fun combine21894(left: Int, right: Int): Pair21894<Int, Int> {
    return Pair21894(left + 69, right + 65)
}
sealed class Outcome21895 {
    data class Success(val payload: String) : Outcome21895()
    data object Failure : Outcome21895()
}

fun describe21895(outcome: Outcome21895): String = when (outcome) {
    is Outcome21895.Success -> outcome.payload
    Outcome21895.Failure -> "failed"
}
@JvmName("step21896")
fun step21896(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
sealed class Outcome21897 {
    data class Success(val payload: String) : Outcome21897()
    data object Failure : Outcome21897()
}

fun describe21897(outcome: Outcome21897): String = when (outcome) {
    is Outcome21897.Success -> outcome.payload
    Outcome21897.Failure -> "failed"
}
sealed class Outcome21898 {
    data class Success(val payload: String) : Outcome21898()
    data object Failure : Outcome21898()
}

fun describe21898(outcome: Outcome21898): String = when (outcome) {
    is Outcome21898.Success -> outcome.payload
    Outcome21898.Failure -> "failed"
}
