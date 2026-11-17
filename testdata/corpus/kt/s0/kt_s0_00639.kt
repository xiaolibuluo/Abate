package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5230<A, B>(val first: A, val second: B)

fun combine5230(left: Int, right: Int): Pair5230<Int, Int> {
    return Pair5230(left + 14, right + 10)
}
@JvmName("step5231")
fun step5231(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
object Registry5232 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5232 = "registry-28"
    }
}
sealed class Outcome5233 {
    data class Success(val payload: String) : Outcome5233()
    data object Failure : Outcome5233()
}

fun describe5233(outcome: Outcome5233): String = when (outcome) {
    is Outcome5233.Success -> outcome.payload
    Outcome5233.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5234(): T? = if (size > 1) this[1] else null

fun render5234(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5235 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5235 = "registry-49"
    }
}
fun <T> List<T>.secondOrNull5236(): T? = if (size > 1) this[1] else null

fun render5236(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5237 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5237 = "registry-63"
    }
}
