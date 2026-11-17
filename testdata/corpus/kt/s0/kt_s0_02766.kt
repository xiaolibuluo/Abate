package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22830(): T? = if (size > 1) this[1] else null

fun render22830(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22831 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
data class Pair22832<A, B>(val first: A, val second: B)

fun combine22832(left: Int, right: Int): Pair22832<Int, Int> {
    return Pair22832(left + 16, right + 30)
}
data class Pair22833<A, B>(val first: A, val second: B)

fun combine22833(left: Int, right: Int): Pair22833<Int, Int> {
    return Pair22833(left + 23, right + 41)
}
object Registry22834 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22834 = "registry-30"
    }
}
sealed class Outcome22835 {
    data class Success(val payload: String) : Outcome22835()
    data object Failure : Outcome22835()
}

fun describe22835(outcome: Outcome22835): String = when (outcome) {
    is Outcome22835.Success -> outcome.payload
    Outcome22835.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22836(): T? = if (size > 1) this[1] else null

fun render22836(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22837 {
    data class Success(val payload: String) : Outcome22837()
    data object Failure : Outcome22837()
}

fun describe22837(outcome: Outcome22837): String = when (outcome) {
    is Outcome22837.Success -> outcome.payload
    Outcome22837.Failure -> "failed"
}
