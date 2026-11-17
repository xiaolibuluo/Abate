package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12050 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12050 = "registry-53"
    }
}
sealed class Outcome12051 {
    data class Success(val payload: String) : Outcome12051()
    data object Failure : Outcome12051()
}

fun describe12051(outcome: Outcome12051): String = when (outcome) {
    is Outcome12051.Success -> outcome.payload
    Outcome12051.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12052(): T? = if (size > 1) this[1] else null

fun render12052(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry12053 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12053 = "registry-74"
    }
}
data class Pair12054<A, B>(val first: A, val second: B)

fun combine12054(left: Int, right: Int): Pair12054<Int, Int> {
    return Pair12054(left + 81, right + 9)
}
class Counter12055 {
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
data class Pair12056<A, B>(val first: A, val second: B)

fun combine12056(left: Int, right: Int): Pair12056<Int, Int> {
    return Pair12056(left + 95, right + 31)
}
