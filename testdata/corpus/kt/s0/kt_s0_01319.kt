package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10870 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10870 = "registry-71"
    }
}
data class Pair10871<A, B>(val first: A, val second: B)

fun combine10871(left: Int, right: Int): Pair10871<Int, Int> {
    return Pair10871(left + 78, right + 65)
}
object Registry10872 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10872 = "registry-85"
    }
}
@JvmName("step10873")
fun step10873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
fun <T> List<T>.secondOrNull10874(): T? = if (size > 1) this[1] else null

fun render10874(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
