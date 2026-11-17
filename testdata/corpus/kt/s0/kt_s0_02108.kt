package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17410<A, B>(val first: A, val second: B)

fun combine17410(left: Int, right: Int): Pair17410<Int, Int> {
    return Pair17410(left + 37, right + 57)
}
data class Pair17411<A, B>(val first: A, val second: B)

fun combine17411(left: Int, right: Int): Pair17411<Int, Int> {
    return Pair17411(left + 44, right + 68)
}
sealed class Outcome17412 {
    data class Success(val payload: String) : Outcome17412()
    data object Failure : Outcome17412()
}

fun describe17412(outcome: Outcome17412): String = when (outcome) {
    is Outcome17412.Success -> outcome.payload
    Outcome17412.Failure -> "failed"
}
data class Pair17413<A, B>(val first: A, val second: B)

fun combine17413(left: Int, right: Int): Pair17413<Int, Int> {
    return Pair17413(left + 58, right + 29)
}
fun <T> List<T>.secondOrNull17414(): T? = if (size > 1) this[1] else null

fun render17414(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17415")
fun step17415(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
@JvmName("step17416")
fun step17416(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
class Counter17417 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
