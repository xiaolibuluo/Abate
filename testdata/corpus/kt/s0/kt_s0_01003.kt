package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8220(): T? = if (size > 1) this[1] else null

fun render8220(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8221<A, B>(val first: A, val second: B)

fun combine8221(left: Int, right: Int): Pair8221<Int, Int> {
    return Pair8221(left + 28, right + 40)
}
data class Pair8222<A, B>(val first: A, val second: B)

fun combine8222(left: Int, right: Int): Pair8222<Int, Int> {
    return Pair8222(left + 35, right + 51)
}
fun <T> List<T>.secondOrNull8223(): T? = if (size > 1) this[1] else null

fun render8223(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8224 {
    var current: Int by Delegates.observable(49) { _, _, new -> println(new) }

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
sealed class Outcome8225 {
    data class Success(val payload: String) : Outcome8225()
    data object Failure : Outcome8225()
}

fun describe8225(outcome: Outcome8225): String = when (outcome) {
    is Outcome8225.Success -> outcome.payload
    Outcome8225.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8226(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8227 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8227 = "registry-70"
    }
}
