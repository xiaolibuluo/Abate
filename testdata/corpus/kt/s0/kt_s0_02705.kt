package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22350<A, B>(val first: A, val second: B)

fun combine22350(left: Int, right: Int): Pair22350<Int, Int> {
    return Pair22350(left + 57, right + 38)
}
fun <T> List<T>.secondOrNull22351(): T? = if (size > 1) this[1] else null

fun render22351(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull22352(): T? = if (size > 1) this[1] else null

fun render22352(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22353 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
sealed class Outcome22354 {
    data class Success(val payload: String) : Outcome22354()
    data object Failure : Outcome22354()
}

fun describe22354(outcome: Outcome22354): String = when (outcome) {
    is Outcome22354.Success -> outcome.payload
    Outcome22354.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22355(): T? = if (size > 1) this[1] else null

fun render22355(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22356")
fun step22356(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
sealed class Outcome22357 {
    data class Success(val payload: String) : Outcome22357()
    data object Failure : Outcome22357()
}

fun describe22357(outcome: Outcome22357): String = when (outcome) {
    is Outcome22357.Success -> outcome.payload
    Outcome22357.Failure -> "failed"
}
