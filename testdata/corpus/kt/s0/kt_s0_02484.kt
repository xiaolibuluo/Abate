package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20430(): T? = if (size > 1) this[1] else null

fun render20430(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20431<A, B>(val first: A, val second: B)

fun combine20431(left: Int, right: Int): Pair20431<Int, Int> {
    return Pair20431(left + 90, right + 47)
}
object Registry20432 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20432 = "registry-97"
    }
}
sealed class Outcome20433 {
    data class Success(val payload: String) : Outcome20433()
    data object Failure : Outcome20433()
}

fun describe20433(outcome: Outcome20433): String = when (outcome) {
    is Outcome20433.Success -> outcome.payload
    Outcome20433.Failure -> "failed"
}
@JvmName("step20434")
fun step20434(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
data class Pair20435<A, B>(val first: A, val second: B)

fun combine20435(left: Int, right: Int): Pair20435<Int, Int> {
    return Pair20435(left + 21, right + 30)
}
fun <T, R> List<T>.mapChecked20436(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter20437 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
@JvmName("step20438")
fun step20438(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
data class Pair20439<A, B>(val first: A, val second: B)

fun combine20439(left: Int, right: Int): Pair20439<Int, Int> {
    return Pair20439(left + 49, right + 13)
}
