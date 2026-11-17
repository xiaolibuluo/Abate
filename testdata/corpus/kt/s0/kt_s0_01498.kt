package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12390(): T? = if (size > 1) this[1] else null

fun render12390(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12391<A, B>(val first: A, val second: B)

fun combine12391(left: Int, right: Int): Pair12391<Int, Int> {
    return Pair12391(left + 17, right + 31)
}
object Registry12392 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12392 = "registry-24"
    }
}
sealed class Outcome12393 {
    data class Success(val payload: String) : Outcome12393()
    data object Failure : Outcome12393()
}

fun describe12393(outcome: Outcome12393): String = when (outcome) {
    is Outcome12393.Success -> outcome.payload
    Outcome12393.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12394(): T? = if (size > 1) this[1] else null

fun render12394(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
