package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20880(): T? = if (size > 1) this[1] else null

fun render20880(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20881")
fun step20881(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
fun <T, R> List<T>.mapChecked20882(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step20883")
fun step20883(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
data class Pair20884<A, B>(val first: A, val second: B)

fun combine20884(left: Int, right: Int): Pair20884<Int, Int> {
    return Pair20884(left + 17, right + 29)
}
class Counter20885 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull20886(): T? = if (size > 1) this[1] else null

fun render20886(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20887(): T? = if (size > 1) this[1] else null

fun render20887(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome20888 {
    data class Success(val payload: String) : Outcome20888()
    data object Failure : Outcome20888()
}

fun describe20888(outcome: Outcome20888): String = when (outcome) {
    is Outcome20888.Success -> outcome.payload
    Outcome20888.Failure -> "failed"
}
