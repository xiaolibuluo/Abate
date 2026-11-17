package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7720<A, B>(val first: A, val second: B)

fun combine7720(left: Int, right: Int): Pair7720<Int, Int> {
    return Pair7720(left + 50, right + 45)
}
class Counter7721 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7722(): T? = if (size > 1) this[1] else null

fun render7722(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7723<A, B>(val first: A, val second: B)

fun combine7723(left: Int, right: Int): Pair7723<Int, Int> {
    return Pair7723(left + 71, right + 17)
}
data class Pair7724<A, B>(val first: A, val second: B)

fun combine7724(left: Int, right: Int): Pair7724<Int, Int> {
    return Pair7724(left + 78, right + 28)
}
@JvmName("step7725")
fun step7725(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
data class Pair7726<A, B>(val first: A, val second: B)

fun combine7726(left: Int, right: Int): Pair7726<Int, Int> {
    return Pair7726(left + 92, right + 50)
}
sealed class Outcome7727 {
    data class Success(val payload: String) : Outcome7727()
    data object Failure : Outcome7727()
}

fun describe7727(outcome: Outcome7727): String = when (outcome) {
    is Outcome7727.Success -> outcome.payload
    Outcome7727.Failure -> "failed"
}
sealed class Outcome7728 {
    data class Success(val payload: String) : Outcome7728()
    data object Failure : Outcome7728()
}

fun describe7728(outcome: Outcome7728): String = when (outcome) {
    is Outcome7728.Success -> outcome.payload
    Outcome7728.Failure -> "failed"
}
