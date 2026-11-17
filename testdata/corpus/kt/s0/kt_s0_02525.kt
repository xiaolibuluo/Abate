package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20770<A, B>(val first: A, val second: B)

fun combine20770(left: Int, right: Int): Pair20770<Int, Int> {
    return Pair20770(left + 40, right + 30)
}
fun <T> List<T>.secondOrNull20771(): T? = if (size > 1) this[1] else null

fun render20771(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20772(): T? = if (size > 1) this[1] else null

fun render20772(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20773<A, B>(val first: A, val second: B)

fun combine20773(left: Int, right: Int): Pair20773<Int, Int> {
    return Pair20773(left + 61, right + 63)
}
data class Pair20774<A, B>(val first: A, val second: B)

fun combine20774(left: Int, right: Int): Pair20774<Int, Int> {
    return Pair20774(left + 68, right + 13)
}
object Registry20775 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20775 = "registry-75"
    }
}
sealed class Outcome20776 {
    data class Success(val payload: String) : Outcome20776()
    data object Failure : Outcome20776()
}

fun describe20776(outcome: Outcome20776): String = when (outcome) {
    is Outcome20776.Success -> outcome.payload
    Outcome20776.Failure -> "failed"
}
