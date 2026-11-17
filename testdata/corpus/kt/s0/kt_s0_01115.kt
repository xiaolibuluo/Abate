package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9140")
fun step9140(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
@JvmName("step9141")
fun step9141(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
sealed class Outcome9142 {
    data class Success(val payload: String) : Outcome9142()
    data object Failure : Outcome9142()
}

fun describe9142(outcome: Outcome9142): String = when (outcome) {
    is Outcome9142.Success -> outcome.payload
    Outcome9142.Failure -> "failed"
}
@JvmName("step9143")
fun step9143(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
data class Pair9144<A, B>(val first: A, val second: B)

fun combine9144(left: Int, right: Int): Pair9144<Int, Int> {
    return Pair9144(left + 81, right + 46)
}
class Counter9145 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull9146(): T? = if (size > 1) this[1] else null

fun render9146(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair9147<A, B>(val first: A, val second: B)

fun combine9147(left: Int, right: Int): Pair9147<Int, Int> {
    return Pair9147(left + 5, right + 18)
}
