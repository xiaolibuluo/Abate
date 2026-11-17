package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13610<A, B>(val first: A, val second: B)

fun combine13610(left: Int, right: Int): Pair13610<Int, Int> {
    return Pair13610(left + 44, right + 20)
}
sealed class Outcome13611 {
    data class Success(val payload: String) : Outcome13611()
    data object Failure : Outcome13611()
}

fun describe13611(outcome: Outcome13611): String = when (outcome) {
    is Outcome13611.Success -> outcome.payload
    Outcome13611.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13612(): T? = if (size > 1) this[1] else null

fun render13612(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13613<A, B>(val first: A, val second: B)

fun combine13613(left: Int, right: Int): Pair13613<Int, Int> {
    return Pair13613(left + 65, right + 53)
}
object Registry13614 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13614 = "registry-72"
    }
}
fun <T, R> List<T>.mapChecked13615(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair13616<A, B>(val first: A, val second: B)

fun combine13616(left: Int, right: Int): Pair13616<Int, Int> {
    return Pair13616(left + 86, right + 25)
}
fun <T, R> List<T>.mapChecked13617(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13618 {
    data class Success(val payload: String) : Outcome13618()
    data object Failure : Outcome13618()
}

fun describe13618(outcome: Outcome13618): String = when (outcome) {
    is Outcome13618.Success -> outcome.payload
    Outcome13618.Failure -> "failed"
}
class Counter13619 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
