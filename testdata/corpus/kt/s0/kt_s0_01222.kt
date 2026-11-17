package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10090<A, B>(val first: A, val second: B)

fun combine10090(left: Int, right: Int): Pair10090<Int, Int> {
    return Pair10090(left + 27, right + 57)
}
sealed class Outcome10091 {
    data class Success(val payload: String) : Outcome10091()
    data object Failure : Outcome10091()
}

fun describe10091(outcome: Outcome10091): String = when (outcome) {
    is Outcome10091.Success -> outcome.payload
    Outcome10091.Failure -> "failed"
}
data class Pair10092<A, B>(val first: A, val second: B)

fun combine10092(left: Int, right: Int): Pair10092<Int, Int> {
    return Pair10092(left + 41, right + 18)
}
@JvmName("step10093")
fun step10093(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
@JvmName("step10094")
fun step10094(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
class Counter10095 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull10096(): T? = if (size > 1) this[1] else null

fun render10096(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
