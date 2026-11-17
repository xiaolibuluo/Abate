package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5300 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5300 = "registry-8"
    }
}
object Registry5301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5301 = "registry-15"
    }
}
sealed class Outcome5302 {
    data class Success(val payload: String) : Outcome5302()
    data object Failure : Outcome5302()
}

fun describe5302(outcome: Outcome5302): String = when (outcome) {
    is Outcome5302.Success -> outcome.payload
    Outcome5302.Failure -> "failed"
}
sealed class Outcome5303 {
    data class Success(val payload: String) : Outcome5303()
    data object Failure : Outcome5303()
}

fun describe5303(outcome: Outcome5303): String = when (outcome) {
    is Outcome5303.Success -> outcome.payload
    Outcome5303.Failure -> "failed"
}
data class Pair5304<A, B>(val first: A, val second: B)

fun combine5304(left: Int, right: Int): Pair5304<Int, Int> {
    return Pair5304(left + 36, right + 42)
}
fun <T> List<T>.secondOrNull5305(): T? = if (size > 1) this[1] else null

fun render5305(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5306 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5306 = "registry-50"
    }
}
object Registry5307 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5307 = "registry-57"
    }
}
@JvmName("step5308")
fun step5308(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
