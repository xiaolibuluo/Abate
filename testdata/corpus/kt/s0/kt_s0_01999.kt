package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16530")
fun step16530(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
@JvmName("step16531")
fun step16531(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
class Counter16532 {
    var current: Int by Delegates.observable(71) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked16533(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16534 {
    var current: Int by Delegates.observable(85) { _, _, new -> println(new) }

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
data class Pair16535<A, B>(val first: A, val second: B)

fun combine16535(left: Int, right: Int): Pair16535<Int, Int> {
    return Pair16535(left + 92, right + 45)
}
data class Pair16536<A, B>(val first: A, val second: B)

fun combine16536(left: Int, right: Int): Pair16536<Int, Int> {
    return Pair16536(left + 99, right + 56)
}
fun <T> List<T>.secondOrNull16537(): T? = if (size > 1) this[1] else null

fun render16537(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16538 {
    data class Success(val payload: String) : Outcome16538()
    data object Failure : Outcome16538()
}

fun describe16538(outcome: Outcome16538): String = when (outcome) {
    is Outcome16538.Success -> outcome.payload
    Outcome16538.Failure -> "failed"
}
